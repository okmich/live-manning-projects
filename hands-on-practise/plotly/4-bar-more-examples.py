import os
import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

from plotly.subplots import make_subplots

## data preparation

vgsales_data = pd.read_csv('datasets/vgsales.csv')

## data preparation complete

subplot_titles = []

# year and count of vgsales
ydf = pd.DataFrame(vgsales_data[['Year']].value_counts()).sort_index()
ydf.reset_index(inplace=True)
ydf.columns = ['Year', 'Count']

year_count_sales = go.Bar(y=ydf.Count, x=ydf.Year)
subplot_titles.append("Number of games sold per Year")
year_count_sales_h = go.Bar(x=ydf.Count, y=ydf.Year, orientation='h')
subplot_titles.append("Number of games sold per Year")


# genre and count of vgsales
gdf = pd.DataFrame(vgsales_data[['Genre']].value_counts()).sort_index()
gdf.reset_index(inplace=True)
gdf.columns = ['Genre', 'Count']
genre_count_sales = go.Bar(y=gdf.Count, x=gdf.Genre)
subplot_titles.append("Number of games sold per Genre")

# Sales by Year and Genre 
sales_by_genre = vgsales_data.groupby('Genre').sum()
na_sales_trace = go.Bar(y=sales_by_genre.NA_Sales, x=sales_by_genre.index, name="NA Sales")
eu_sales_trace = go.Bar(y=sales_by_genre.EU_Sales, x=sales_by_genre.index, name="EU Sales")
jp_sales_trace = go.Bar(y=sales_by_genre.JP_Sales, x=sales_by_genre.index, name="JP Sales")
other_sales_trace = go.Bar(y=sales_by_genre.Other_Sales, x=sales_by_genre.index, name="Other Sales")
global_sales_trace = go.Bar(y=sales_by_genre.Global_Sales, x=sales_by_genre.index, name="Global Sales")

subplot_titles.append("Sales of games per year and Genre")


fig = make_subplots(rows=2, cols=2, subplot_titles=subplot_titles,  vertical_spacing=0.15, horizontal_spacing=0.1)
fig.update_layout(showlegend=False, title_text="Video Game Sales Analysis")


fig.add_trace(year_count_sales, row=1, col=1)
fig.add_trace(year_count_sales_h, row=1, col=2)
fig.add_trace(genre_count_sales, row=2, col=1)

fig.add_trace(na_sales_trace, row=2, col=2)
fig.add_trace(eu_sales_trace, row=2, col=2)
fig.add_trace(jp_sales_trace, row=2, col=2)
fig.add_trace(other_sales_trace, row=2, col=2)
fig.add_trace(global_sales_trace, row=2, col=2)

fig.update_layout(barmode="stack")

# Genre sales

# Genre sales by publisher

# Genre sales by regions 

# Genre sales faceted by Year for the last 6 years



fig.write_html("4-bar.html", auto_open=True)