import numpy as np
import plotly.figure_factory as ff
import plotly.graph_objects as go

import pandas as pd

## data preparation
vgsales_data = pd.read_csv('datasets/vgsales.csv')
## data preparation complete

fig =  ff.create_table(vgsales_data.describe())

for col_name in ['NA_Sales', 'EU_Sales', 'JP_Sales', 'Other_Sales']:
	trace2 = go.Violin(
		y=np.log(vgsales_data[col_name]),
		name=col_name,
		box_visible=True,
		points='outliers',
		meanline_visible=True,
		x0=col_name[:2],
		xaxis='x2', yaxis='y2')
	fig.add_trace(trace2)

fig.update_layout(
    title_text = "Video games sales data",
    margin = {'t':50, 'b':100},
    xaxis = {'domain': [0, .5]},
    xaxis2 = {'domain': [0.6, 1.], 'title':"Regions"},
    yaxis2 = {'anchor': 'x2', 'title': 'log(Sales)'}
)

fig.write_html("12-table-with-charts.html", auto_open=True)