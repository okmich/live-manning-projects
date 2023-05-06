import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

from plotly.subplots import make_subplots

## data preparation
vgsales_data = pd.read_csv('datasets/vgsales.csv')
## data preparation complete

fig = px.scatter(vgsales_data, x="NA_Sales", y="Global_Sales",
						size="EU_Sales",
						color='Genre')

fig.write_html("hello-line-scatter.html", auto_open=True)