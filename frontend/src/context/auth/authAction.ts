// AuthContext Actions

import { UserLoginResponse } from "../../services/auth";

export enum AuthActionEnum {
  LOG_IN = "LOG_IN",
  LOG_OUT = "LOG_OUT",
}

export type AuthAction =
  | {
      type: AuthActionEnum.LOG_IN;
      payload: UserLoginResponse;
    }
  | {
      type: AuthActionEnum.LOG_OUT;
      payload: null;
    };
