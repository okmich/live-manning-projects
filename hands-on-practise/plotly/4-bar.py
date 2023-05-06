import os
import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

## data preparation

data_file = os.path.join(os.getcwd(), "datasets/online-retail/Online Retail.xlsx")
data = pd.DataFrame((pd.read_excel(data_file))[['Country']].value_counts())
data.reset_index(inplace=True)
data.columns = ['Country', 'Count']

## data preparation complete

fig = px.bar(data, x='Country', y='Count', title="Customer Count per Country", text='Count')
fig.update_layout(go.Layout(height=600, width=800))
fig.write_html("hello-bar-0.html", auto_open=True)