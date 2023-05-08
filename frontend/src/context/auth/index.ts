import { createContext } from "react";
import { AuthAction } from "./authAction";
import { AuthState, defaultAuthState } from "./authReducer";

export interface AuthContext {
  authState: AuthState;
  authDispatch: (action: AuthAction) => void;
}

export const authContext = createContext<AuthContext>({
  authState: defaultAuthState,
  // eslint-disable-next-line @typescript-eslint/no-empty-function
  authDispatch: (): void => {},
});
