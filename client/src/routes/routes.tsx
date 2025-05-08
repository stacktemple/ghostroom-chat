import { createBrowserRouter } from "react-router-dom";
import Home from "../pages/Home";
import Layout from "../layouts/Layout";
import Chatroom from "../pages/ChatRoom";
import NotFound from "../pages/NotFound";
import Test from "../pages/Test";
import Unauthorized from "../pages/Unauthorized";
import ProtectedRoomRoute from "./ProtectedRoomRoute";

const routes = createBrowserRouter([
  {
    path: "/",
    element: <Layout />,
    children: [
      { index: true, element: <Home /> },
      {
        path: "room",
        element: <ProtectedRoomRoute />,
        children: [{ path: ":roomName", element: <Chatroom /> }],
      },
      { path: "unauth", element: <Unauthorized /> },
      { path: "*", element: <NotFound /> },
    ],
  },
  {
    path: "test",
    element: <Test />,
  },
]);

export default routes;
