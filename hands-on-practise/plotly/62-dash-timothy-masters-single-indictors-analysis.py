from datetime import datetime

import pandas as pd

from dash import Dash, html, dcc, Input, Output
import plotly.express as px
import plotly.graph_objects as go
from plotly.subplots import make_subplots

def parse_date_column(s):
	return datetime.strptime(s, '%Y%m%d')

app = Dash(__name__)

symbols = ['AUDUSD', 'BTCUSD', 'EURUSD', 'EURGBP', 'USDCHF', 'WTI', 'XAUUSD']
columns = ['date', 'open', 'high', 'low', 'close', 'volume']

all_sym_data = {}
for sym in symbols:
	price_data = pd.read_csv(f'datasets/indicators/{sym}.csv', header=None, index_col=0, 
		names=columns, delimiter=' ', date_parser=parse_date_column)
	out_var_data = pd.read_csv(f'datasets/indicators/{sym}.txt', delimiter='\t', 
		index_col=0, date_parser=parse_date_column)
	all_sym_data[sym] = price_data.join(out_var_data, how='inner')

max_date = all_sym_data[symbols[0]].iloc[-1:].index.to_list()[0]
min_date = all_sym_data[symbols[0]].iloc[0:].index.to_list()[0]

# price_data = raw_df.loc[(raw_df.index >= min_date) & (raw_df.index <= max_date)]

indicators = list(all_sym_data[symbols[0]].columns.to_list())
indicators.sort()

app.layout = html.Div([
		html.H2(children="Analyzing Timothy Master's Indicators", style={'textAlign': 'center'}),
		html.Div(children=[
			html.Div(children=[
			        html.Label('Instrument', style={'text-weight': 'bold'}),
			        dcc.Dropdown(symbols, symbols[0], id='i_symbols'),

			        html.Br(),
			        html.Label('Indicators', style={'text-weight': 'bold'}),
			        dcc.Dropdown(indicators, indicators[0], multi=True, id='i_indicators'),

			        html.Br(),
			        html.Label('Date range', style={'text-weight': 'bold'}),
			        html.Br(),
			        dcc.DatePickerRange(
			            id='i_date_picker_range',
			            start_date=min_date.date(),
			            end_date=max_date.date()
			        ),
			        html.Br(),
			        html.Hr(),
				], style={'padding': 5, 'flex': 1}
			),
			html.Div(children=[
				dcc.Graph(id='display-graph')
				], style={'padding': 5, 'flex': 3, 'float':'right', 'border-left' : '1px solid black'}
			),
			html.Div(children=[
				dcc.Graph(id='display-dist-graph')
				], style={'flex': 1.5, 'float':'right', 'border-left' : '1px solid #ccc'}
			)
		], style={'display': 'flex', 'flex-direction': 'row'}),

	]
)


@app.callback(
	Output('display-graph', 'figure'),
    Input('i_symbols', 'value'),
    Input('i_indicators', 'value'),
    Input('i_date_picker_range', 'start_date'),
    Input('i_date_picker_range', 'end_date')
)
def update_graph(symbol, selected_indicators, start_date, end_date):
	# get the dataframe
	all_data = all_sym_data[symbol]	

	# apply filter
	display_df = all_data.loc[(all_data.index >= start_date) & (all_data.index <= end_date)]
	# ensure the selected_indicator is always a list
	indicator_list = selected_indicators if isinstance(selected_indicators, list) else [selected_indicators]

	no_indicators = len(indicator_list)
	row_heights = [0.2]
	row_heights.extend([0.1 for i in range(no_indicators)])

	fig = make_subplots(
		rows=no_indicators+1, 
		cols=1, shared_xaxes=True,
		vertical_spacing=0.01,
		x_title='Date',
		row_heights=row_heights
	)
	height_guess = 500 + (no_indicators*100)
	
	fig.add_trace(
		go.Candlestick(x=display_df.index,
			open=display_df['open'],
            high=display_df['high'],
            low=display_df['low'],
            close=display_df['close'],
            name="Price data"), 
		row=1, 
		col=1
	)
	fig['layout']['yaxis']['title']="Close Price"
	fig['layout']['yaxis']['title']['font_size'] = 12

	def update_yaxis_lable(yaxis, anchor, label):
		# print(yaxis, anchor, label)
		if yaxis['anchor'] == anchor:
			yaxis['title'] = label
			yaxis['title']['font_size'] = 12

	for i, ind_name in enumerate(indicator_list):
		fig.add_trace(
			go.Scatter(
				x=display_df.index, 
				y=display_df[ind_name],
				name=ind_name), 
			row=i+2, col=1)
		fig.for_each_yaxis(lambda yx : update_yaxis_lable(yx, f'x{i+2}', ind_name))

	fig.update_layout(
		height=height_guess, 
		xaxis_rangeslider_visible=False,
		title_text='Price vs Indicators',
		showlegend=False)

	# re-evaluate min and max dates from the data
	max_date = all_data.iloc[-1:].index.to_list()[0]
	min_date = all_data.iloc[0:].index.to_list()[0]

	return fig


@app.callback(
	Output('display-dist-graph', 'figure'),
    Input('i_symbols', 'value'),
    Input('i_indicators', 'value'),
    Input('i_date_picker_range', 'start_date'),
    Input('i_date_picker_range', 'end_date')
)
def update_dist_graph(symbol, selected_indicators, start_date, end_date):	
	# get the dataframe
	all_data = all_sym_data[symbol]	
	
	display_df = all_data.loc[(all_data.index >= start_date) & (all_data.index <= end_date)]
	# ensure the selected_indicator is always a list
	indicator_list = selected_indicators if isinstance(selected_indicators, list) else [selected_indicators]

	no_indicators = len(indicator_list)
	
	traces = [
		go.Histogram(
			x=display_df[indx], 
			histnorm='probability density',
			nbinsx=20,
			name=indx) for indx in indicator_list
	]

	fig = go.Figure(data=traces)

	fig.update_layout(
		barmode='overlay',
		title_text='Probability distribution',
		showlegend=True,
		bargap=0.1,
		legend_orientation='h')

	return fig


app.run_server(debug=True)