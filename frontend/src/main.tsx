import React from "react";
import ReactDOM from "react-dom/client";
import App from "@/App.tsx";
import "@/index.css";
import { QueryClient, QueryClientProvider } from "react-query";
import { RouterProvider, createBrowserRouter } from "react-router-dom";
import ErrorPage from "@/routes/ErrorPage.tsx";
import PasswordCards from "./routes/PasswordCards.tsx";

const client = new QueryClient({
  defaultOptions: {
    queries: { retryOnMount: false, refetchOnWindowFocus: false },
  },
});
const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [{ path: "", element: <PasswordCards /> }],
  },
]);

ReactDOM.createRoot(document.getElementById("root") as HTMLElement).render(
  <React.StrictMode>
    <QueryClientProvider client={client}>
      <RouterProvider router={router} />
    </QueryClientProvider>
  </React.StrictMode>
);
