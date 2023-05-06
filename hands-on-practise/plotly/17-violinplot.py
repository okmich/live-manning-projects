import numpy as np
import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

from plotly.subplots import make_subplots

## data preparation
vgsales_data = pd.read_csv('datasets/vgsales.csv')
## data preparation complete


fig = make_subplots(rows=1, cols=2, 
	subplot_titles=['Sales distribution across regions', 'Sales log distribution across regions'])

for col_name in ['NA_Sales', 'EU_Sales', 'JP_Sales', 'Other_Sales']:
	trace1 = go.Violin(
		y=vgsales_data[col_name],
		name=col_name,
		box_visible=True,
		points='outliers',
		meanline_visible=True,
		x0=col_name)
	fig.add_trace(trace1, row=1, col=1)


for col_name in ['NA_Sales', 'EU_Sales', 'JP_Sales', 'Other_Sales']:
	trace2 = go.Violin(
		y=np.log(vgsales_data[col_name]),
		name=col_name,
		box_visible=True,
		points='outliers',
		meanline_visible=True,
		x0=col_name)
	fig.add_trace(trace2, row=1, col=2)


fig.update_traces(box_visible=True, meanline_visible=True)
fig.update_layout(violinmode='group')
fig.write_html('17-violin.html', auto_open=True)


# more here https://plotly.com/python/violin/