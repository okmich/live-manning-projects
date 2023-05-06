from dash import Dash, html, dcc
from datetime import date

app = Dash(__name__)

app.layout = html.Div([
    html.Div(children=[
        html.Label('Dropdown'),
        dcc.Dropdown(['New York City', 'Montréal', 'San Francisco'], 'Montréal'),

        html.Br(),
        html.Label('Multi-Select Dropdown'),
        dcc.Dropdown(['New York City', 'Montréal', 'San Francisco'],
                     ['Montréal', 'San Francisco'],
                     multi=True),

        html.Br(),
        html.Label('Radio Items'),
        dcc.RadioItems(['New York City', 'Montréal', 'San Francisco'], 'Montréal'),

        html.Br(),
        html.Label('Date input : '),
        dcc.DatePickerSingle(
            id='date-picker-single',
            date=date(1997, 5, 10)
        ),

        html.Br(),
        html.Label('Date range input: '),
        dcc.DatePickerRange(
            id='date-picker-range',
            start_date=date(1997, 5, 3),
            end_date_placeholder_text='Select a date!'
        ),

        html.Br(),
        html.Br(),
        html.Label('Text input: '),
        dcc.Input(placeholder='Enter a value...', type='text', value=''),

        html.Br(),
        html.Label('Text area: '),
        dcc.Textarea(
            placeholder='Enter a value...',
            value='This is a TextArea component',
            style={'width': '100%'}
        ),

        html.Br(),
        html.Button('Submit Button', id='button-example-1', style= {'float':'right'}),

        html.Br(),
        html.Label('Graph: '),
        dcc.Graph(
            figure=dict(
                data=[
                    dict(
                        x=[1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003,
                        2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012],
                        y=[219, 146, 112, 127, 124, 180, 236, 207, 236, 263,
                        350, 430, 474, 526, 488, 537, 500, 439],
                        name='Rest of world',
                        marker=dict(
                            color='rgb(55, 83, 109)'
                        )
                    ),
                    dict(
                        x=[1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003,
                        2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012],
                        y=[16, 13, 10, 11, 28, 37, 43, 55, 56, 88, 105, 156, 270,
                        299, 340, 403, 549, 499],
                        name='China',
                        marker=dict(
                            color='rgb(26, 118, 255)'
                        )
                    )
                ],
                layout=dict(
                    title='US Export of Plastic Scrap',
                    showlegend=True,
                    margin=dict(l=40, r=0, t=40, b=30)
                )
            ),
            style={'height': 300},
            id='my-graph-example'
        )
    ], style={'padding': 10, 'flex': 1}),

    html.Div(children=[
        html.Label('Checkboxes'),
        dcc.Checklist(['New York City', 'Montréal', 'San Francisco'],
                      ['Montréal', 'San Francisco']
        ),

        html.Br(),
        html.Label('Text Input'),
        dcc.Input(value='MTL', type='text'),

        html.Br(),
        html.Label('Slider'),
        dcc.Slider(
            min=0,
            max=9,
            marks={i: f'Label {i}' if i == 1 else str(i) for i in range(1, 6)},
            value=5,
        ),

        html.Br(),
        dcc.ConfirmDialog(
            id='confirm',
            message='Danger danger! Are you sure you want to continue?'
        ),

        html.Hr(),

        html.Br(),
        html.Label('Link (or Location): '),
        html.Br(),
        html.A(id='component', href="https://dash.plotly.com/dash-core-components", children='Component reference'),
        html.Br(),
        html.A(id='html', href="https://dash.plotly.com/dash-html-components", children='HTML Dash reference'),

        html.Br(),
        dcc.Store(id='my-store', data={'my-data': 'data'}),


        dcc.Loading([
            # ...
        ])

    ], style={'padding': 10, 'flex': 1})
], style={'display': 'flex', 'flex-direction': 'row'})

app.run_server(debug=True)