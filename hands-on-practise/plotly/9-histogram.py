import numpy as np
import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

from plotly.subplots import make_subplots

## data preparation

vgsales_data = pd.read_csv('datasets/vgsales.csv')

## data preparation complete

sub_titles = []
traces = []


# histogram with ordinal x-axis
platform_distribution_trace = go.Histogram(x=vgsales_data.Platform, 
	histfunc='count')
traces.append((platform_distribution_trace, 1, 1))
sub_titles.append(f"Platform distribution")

genre_distribution_trace = go.Histogram(x=vgsales_data.Genre, 
	histfunc='count')
traces.append((genre_distribution_trace, 1, 2))
sub_titles.append(f"Genre distribution")

genre_sales_distribution_trace = go.Histogram(x=vgsales_data.Genre, 
	y=vgsales_data.Global_Sales, histnorm="probability", histfunc='avg')
traces.append((genre_sales_distribution_trace, 1, 3))
sub_titles.append(f"Genre Sales distribution")


def add_traces(area_code, row_num):
	col_name = f"{area_code}_Sales"
	trace1 = go.Histogram(
		x=vgsales_data[col_name],
		nbinsx=10,
		name=col_name)
	traces.append((trace1, row_num, 1))
	sub_titles.append(f"Sales distribution for {col_name}")

	trace2 = go.Histogram(
		x=np.log(vgsales_data[col_name]),
		nbinsx=10,
		name=col_name)
	traces.append((trace2, row_num, 2))
	sub_titles.append(f"Sales log distribution for {col_name}")

	trace3 = go.Histogram(
		x=np.log(vgsales_data[col_name]),
		nbinsx=10,
		name=col_name,
		histnorm="probability density")
	traces.append((trace3, row_num, 3))
	sub_titles.append(f"Sales probability density distribution for {col_name}")

# the left has a histogram of a sales field probability distribution, the second, hold the log_y=True of the first, the third has a denstiy normalized version
for (index, name) in enumerate(["NA","EU","JP","Other","Global"]):
	add_traces(name, index+2)

# create a histogram in a subplot of 5,3
fig = make_subplots(rows=6, cols=3, subplot_titles=sub_titles)

for t in traces:
	fig.add_trace(t[0], row=t[1], col=t[2])

fig.update_layout(title_text="Sales histograms distribution", bargap=0.1, autosize=True, height=1200)

fig.write_html('9-histogram.html', auto_open=True)