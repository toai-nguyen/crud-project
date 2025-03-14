import { BrowserRouter, Routes, Route  } from "react-router-dom";
import Home from "./Pages/Home";
import HomeLayouts from "./Layouts/HomeLayouts";
import Profile from "./Pages/Profile";
import UserForm from "./Pages/UserForm";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<HomeLayouts />}>
          <Route index element={<Home />} />
        </Route>
        <Route path="user/:id" element={<HomeLayouts />}>
          <Route index element={<Profile />} />
        </Route>
        <Route path="/add" element={<HomeLayouts />}>
          <Route index element={<UserForm />} />
        </Route>
        <Route path="/edit/:id" element={<HomeLayouts />}>
          <Route index element={<UserForm isEdit = {true} />} />
        </Route>
      </Routes>
    </BrowserRouter>
  );
}

export default App;
