import React, { useMemo, useReducer } from "react";

import { authContext } from ".";
import authReducer, { defaultAuthState } from "./authReducer";

type Props = {
  children: React.ReactNode;
};

export const AuthContextProvider: React.FC<Props> = ({ children }) => {
  const [authState, authDispatch] = useReducer(authReducer, defaultAuthState);

  // localstorage にユーザー情報が保存されている場合
  // useEffect(() => {
  //   const user = localStorage.getItem("user");
  //   if (user) {
  //     const userLoginData = JSON.parse(user);
  //     authDispatch({
  //       type: AuthActionEnum.LOG_IN,
  //       payload: userLoginData,
  //     });
  //   }
  // }, []);

  const value = useMemo(
    () => ({
      authState,
      authDispatch,
    }),
    [authState]
  );

  return <authContext.Provider value={value}>{children}</authContext.Provider>;
};
