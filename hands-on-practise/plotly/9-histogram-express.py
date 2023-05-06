import pandas as pd
import plotly.express as px
import plotly.graph_objects as go

## data preparation

udemy_courses = pd.read_csv('datasets/udemy_output_All_Business_p1_p626.csv')

## data preparation complete

fig = px.histogram(udemy_courses, x='avg_rating', histnorm='percent', nbins=20, log_y=True, color='is_wishlisted')

fig.write_html('hello-histogram.html', auto_open=True)