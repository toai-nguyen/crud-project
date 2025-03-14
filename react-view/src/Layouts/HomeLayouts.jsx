import { Outlet } from "react-router-dom";
import { Link } from "react-router-dom";
export default function HomeLayouts({}) {
  return (
    <div className="HomeLayouts">
      <header>
        <h1>CRUD using React and Go</h1>
      </header>
      <div>
      </div>
      <main>
        <Outlet />
      </main>
    </div>
  );
}
