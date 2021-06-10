import React from "react";
import { Container } from 'react-bootstrap';
import { BrowserRouter, Route, Switch } from 'react-router-dom'

import Info from './Info/info';
import QuoteBook from './QuoteBook/quotebook';

function Router() {
    return (
        <div>
            <BrowserRouter>
                <Switch>
                    <Route exact
                        path='/'
                        render={() =>
                            <Container fluid="sm" style={{ paddingTop: "20px" }}>
                                <h1>Online Quote Book</h1>
                            </Container>
                        }
                    />
                    <Route exact
                        path='/info'
                        component={Info}
                    />                    
                    <Route exact
                        path='/quotebooks/:quotebookname'
                        component={QuoteBook}
                    />
                </Switch>
            </BrowserRouter>
        </div>
    );
}

export default Router;
