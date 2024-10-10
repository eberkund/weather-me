import { Route, Switch } from "wouter";
import { Settings } from "./pages/Settings";
import { Index } from "./pages/Index";
import clouds from "~/assets/clouds3.jpg";

export default function App() {
  return (
    <div
      style={{
        backgroundSize: "cover",
        backgroundImage: `url(${clouds})`,
      }}
    >
      <div className="bg-gradient-to-t from-white/0 to-white/40 h-lvh w-full">
        <Switch>
          <Route path="/" component={Index} />
          <Route path="/settings" component={Settings} />
          <Route>404: No such page!</Route>
        </Switch>
      </div>
    </div>
  );
}
/* iPhone 11 Pro Max - 1 */

// position: absolute;
// width: 290px;
// height: 245px;
// left: 0px;
// top: 0px;

// background: radial-gradient(100% 100% at 0% 0%, rgba(255, 255, 255, 0.4) 0%, rgba(255, 255, 255, 0) 100%) /* warning: gradient uses a rotation that is not supported by CSS and may not behave as expected */;
// box-shadow: inset 5px 5px 4px rgba(255, 255, 255, 0.1);
// backdrop-filter: blur(21px);
// /* Note: backdrop-filter has minimal browser support */
// border-radius: 20px;
