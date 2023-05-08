import React, { useMemo, useReducer } from "react";

import { authContext } from ".";
import authReducer, { defaultAuthState } from "./authReducer";

type Props = {
  children: React.ReactNode;
};

export const AuthContextProvider: React.FC<Props> = ({ children }) => {
  const [authState, authDispatch] = useReducer(authReducer, defaultAuthState);

  const value = useMemo(
    () => ({
      authState,
      authDispatch,
    }),
    [authState]
  );

  return <authContext.Provider value={value}>{children}</authContext.Provider>;
};
