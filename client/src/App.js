import React from 'react';
import { Switch, BrowserRouter as Router, Route} from "react-router-dom";
import { ThemeProvider, CSSReset ,DarkMode} from "@chakra-ui/core";
import Landing from "./Components/Landing/Landing";
import StudentView from "./Components/StudentView/StudentView";
import InstructorView from "./Components/InstructorView/InstructorView";

function App() {
  return (
    <ThemeProvider>
      <CSSReset />
      
      <Router>
        <Switch>
          <Route path="/student">
            <StudentView />
          </Route>
          <Route path="/instructor">
            <InstructorView />
          </Route>
          <Route path="/">
            <Landing />
          </Route>
        </Switch>
       </Router>
       
    </ThemeProvider>
  );
}

export default App;
