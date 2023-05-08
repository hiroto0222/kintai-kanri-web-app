// AuthContext Reducer

import { Reducer } from "react";
import { User } from "../../services/auth";
import { AuthAction, AuthActionEnum } from "./authAction";

export interface AuthState {
  isLoggedIn: boolean;
  accessToken?: string;
  refreshToken?: string;
  user?: User;
}

export const defaultAuthState: AuthState = {
  isLoggedIn: false,
};

const authReducer: Reducer<AuthState, AuthAction> = (state, action) => {
  if (action.type === AuthActionEnum.LOG_IN) {
    return {
      ...state,
      isLoggedIn: true,
      accessToken: action.payload.access_token,
      refreshToken: action.payload.refresh_token,
      user: action.payload.user,
    };
  }

  if (action.type === AuthActionEnum.LOG_OUT) {
    return defaultAuthState;
  }

  if (action.type === AuthActionEnum.SET_ACCESS_TOKEN) {
    return {
      ...state,
      accessToken: action.payload,
    };
  }

  return defaultAuthState;
};

export default authReducer;
