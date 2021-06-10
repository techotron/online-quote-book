import React from "react";
import { Container } from 'react-bootstrap';
import { BrowserRouter, Route, Switch } from 'react-router-dom'

import Info from './Info/info';

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
            render={() => 
              <Info />              
            }
          />
        </Switch>
      </BrowserRouter>
    </div>
  );
}

export default Router;
