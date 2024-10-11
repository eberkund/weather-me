import { Route, Switch } from "wouter";
import { Settings } from "~/pages/Settings";
import { Index } from "~/pages/Index";

export default function App() {
  return (
    <Switch>
      <Route path="/" component={Index} />
      <Route path="/settings" component={Settings} />
      <Route>404: No such page!</Route>
    </Switch>
  );
}
