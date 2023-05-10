import { HelmetProvider } from "react-helmet-async";
import { Toaster } from "react-hot-toast";
import { BrowserRouter } from "react-router-dom";
import { AuthContextProvider } from "./context/auth/AuthContextProvider";
import Router from "./routes";
import ThemeProvider from "./theme/index.tsx";

const App = () => {
  return (
    <HelmetProvider>
      <BrowserRouter>
        <AuthContextProvider>
          <ThemeProvider>
            <Toaster />
            <Router />
          </ThemeProvider>
        </AuthContextProvider>
      </BrowserRouter>
    </HelmetProvider>
  );
};

export default App;
