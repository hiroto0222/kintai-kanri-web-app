// AuthContext Reducer

import { Reducer } from "react";
import { Employee } from "../../services/employees";
import { AuthAction, AuthActionEnum } from "./authAction";

export interface AuthState {
  isLoggedIn: boolean;
  accessToken?: string;
  refreshToken?: string;
  user?: Employee;
}

export const defaultAuthState: AuthState = {
  isLoggedIn: false,
};

const authReducer: Reducer<AuthState, AuthAction> = (state, action) => {
  if (action.type === AuthActionEnum.LOG_IN) {
    const authState = {
      ...state,
      isLoggedIn: true,
      accessToken: action.payload.access_token,
      refreshToken: action.payload.refresh_token,
      user: action.payload.user,
    };
    localStorage.setItem("accessToken", authState.accessToken);
    localStorage.setItem("refreshToken", authState.refreshToken);
    return authState;
  }

  if (action.type === AuthActionEnum.LOG_OUT) {
    localStorage.removeItem("accessToken");
    localStorage.removeItem("refreshToken");
    return defaultAuthState;
  }

  if (action.type === AuthActionEnum.SET_ACCESS_TOKEN) {
    const authState = {
      ...state,
      accessToken: action.payload,
    };
    localStorage.setItem("accessToken", authState.accessToken);
    return authState;
  }

  return defaultAuthState;
};

export default authReducer;
