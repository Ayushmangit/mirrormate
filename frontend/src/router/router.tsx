import { createBrowserRouter } from "react-router-dom";

import App from "../App";
import LandingPage from "../pages/PublicPages/LandingPage";
import LoginPage from "../pages/PublicPages/LoginPage";
import RegisterPage from "../pages/PublicPages/RegisterPage";
export const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,

    children: [
      {
        index: true,
        element: <LandingPage />,
      },

      {
        path: "login",
        element: <LoginPage />,
      },

      {
        path: "register",
        element: <RegisterPage />,
      },

      // {
      //   element: <ProtectedRoute />,
      //
      //   children: [
      //     {
      //       path: "boards",
      //       element: <BoardsPage />,
      //     },
      //
      //     {
      //       path: "boards/:boardID",
      //       element: <BoardPage />,
      //     },
      //   ],
      // },
    ],
  },
]);
