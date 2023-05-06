import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

from plotly.subplots import make_subplots

## data preparation
vgsales_data = pd.read_csv('datasets/vgsales.csv')
vgsales_data = vgsales_data[vgsales_data['Year'] != 'Adventure']
## data preparation complete

subplot_titles = []

# Sales by Year and Genre 
sales_by_year = vgsales_data.groupby('Year').sum().sort_index()
na_sales_trace = go.Scatter(y=sales_by_year.NA_Sales, x=sales_by_year.index, name="NA Sales", mode="lines+markers")
eu_sales_trace = go.Scatter(y=sales_by_year.EU_Sales, x=sales_by_year.index, name="EU Sales", mode="lines+markers")
jp_sales_trace = go.Scatter(y=sales_by_year.JP_Sales, x=sales_by_year.index, name="JP Sales", mode="lines+markers")
other_sales_trace = go.Scatter(y=sales_by_year.Other_Sales, x=sales_by_year.index, name="Other Sales", mode="lines+markers")
global_sales_trace = go.Scatter(y=sales_by_year.Global_Sales, x=sales_by_year.index, name="Global Sales", mode="lines+markers")
subplot_titles.append("Sales of games per year")

# Compare Regions to Global Sales
eu_global_compare = go.Scatter(y=vgsales_data.Global_Sales, x=vgsales_data.EU_Sales, mode='markers', marker=dict(symbol='triangle-up'))
jp_global_compare = go.Scatter(y=vgsales_data.Global_Sales, x=vgsales_data.JP_Sales, mode='markers', marker=dict(symbol='hexagon'))
na_global_compare = go.Scatter(y=vgsales_data.Global_Sales, x=vgsales_data.NA_Sales, mode='markers', marker=dict(symbol='square'))
other_global_compare = go.Scatter(y=vgsales_data.Global_Sales, x=vgsales_data.Other_Sales, mode='markers', marker=dict(symbol='star'))
subplot_titles.append("Compare Regions Sales with global sales")


# genre and count of vgsales
gdf = pd.DataFrame(vgsales_data[['Genre']].value_counts()).sort_index()
gdf.reset_index(inplace=True)
gdf.columns = ['Genre', 'Count']
genre_count_sales = go.Bar(y=gdf.Count, x=gdf.Genre)
subplot_titles.append("Number of games sold per Genre")



fig = make_subplots(rows=2, cols=2, vertical_spacing=0.15, horizontal_spacing=0.1,
    specs=[[{"colspan": 2}, None],
           [{}, {}]],
    subplot_titles=subplot_titles)
fig.update_layout(title_text="Video Game Sales Analysis")


fig.add_trace(eu_global_compare, row=2, col=1)
fig.add_trace(jp_global_compare, row=2, col=1)
fig.add_trace(na_global_compare, row=2, col=1)
fig.add_trace(other_global_compare, row=2, col=1)

fig.add_trace(genre_count_sales, row=2, col=2)

fig.add_trace(na_sales_trace, row=1, col=1)
fig.add_trace(eu_sales_trace, row=1, col=1)
fig.add_trace(jp_sales_trace, row=1, col=1)
fig.add_trace(other_sales_trace, row=1, col=1)
fig.add_trace(global_sales_trace, row=1, col=1)

# Genre sales

# Genre sales by publisher

# Genre sales by regions 

# Genre sales faceted by Year for the last 6 years



fig.write_html("5-line-scatter.html", auto_open=True)