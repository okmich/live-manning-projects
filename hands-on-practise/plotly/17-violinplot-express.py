import pandas as pd
import plotly.express as px

from plotly.subplots import make_subplots

## data preparation
orders = pd.read_csv('datasets/orders/orders.csv')
## data preparation complete

fig = px.violin(orders, y="Profit", box=True, points='outliers', color='Category', hover_data=orders.columns)

fig.update_layout(width=800, height=700, title_text="Profit Analysis")
fig.write_html('hello-violin.html', auto_open=True)