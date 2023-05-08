// AuthContext Reducer

import { Reducer } from "react";
import { User } from "../../services/auth";
import { AuthAction, AuthActionEnum } from "./authAction";

export interface AuthState {
  isLoggedIn: boolean;
  accessToken?: string;
  user?: User;
}

export const defaultAuthState: AuthState = {
  isLoggedIn: false,
};

const authReducer: Reducer<AuthState, AuthAction> = (state, action) => {
  if (action.type === AuthActionEnum.LOG_IN) {
    localStorage.setItem("user", JSON.stringify(action.payload));
    return {
      ...state,
      isLoggedIn: true,
      accessToken: action.payload.access_token,
      user: action.payload.user,
    };
  }

  if (action.type === AuthActionEnum.LOG_OUT) {
    localStorage.removeItem("user");
    return defaultAuthState;
  }

  return defaultAuthState;
};

export default authReducer;
