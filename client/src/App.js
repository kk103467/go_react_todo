
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import ViewPage from './pages/ViewPage';
import './App.css';

function App() {
  return (
    <Router>
      <div className="App">
        <Switch>
          <Route exact path="/">
            <ViewPage />
          </Route>
        </Switch>
      </div>
    </Router>
  );
}

export default App;