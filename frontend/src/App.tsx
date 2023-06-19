import { Outlet } from "react-router-dom";
import "./App.css";

function App() {
  return (
    <div className="flex items-center flex-col py-12">
      <div className="flex items-center flex-col gap-5">
        <h1 className="text-3xl text-blue-700">Password Manager</h1>
        <Outlet />
      </div>
    </div>
  );
}

export default App;
