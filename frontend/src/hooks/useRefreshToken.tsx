import { AxiosResponse } from "axios";
import { useContext } from "react";
import { authContext } from "../context/auth";
import { AuthActionEnum } from "../context/auth/authAction";
import api from "../services/api";
import { RefreshAccessTokenResponse } from "../services/auth";

const useRefreshToken = () => {
  const { authState, authDispatch } = useContext(authContext);

  const refresh = async () => {
    const response: AxiosResponse<RefreshAccessTokenResponse> = await api.post(
      "/auth/refresh",
      {
        refresh_token: authState.refreshToken,
      },
      {
        withCredentials: true,
      }
    );

    authDispatch({
      type: AuthActionEnum.SET_ACCESS_TOKEN,
      payload: response.data.access_token,
    });

    return response.data.access_token;
  };

  return refresh;
};

export default useRefreshToken;
