import { useContext, useEffect } from "react";
import { authContext } from "../context/auth";
import { AuthActionEnum } from "../context/auth/authAction";
import axiosInstance from "../services/api";
import useRefreshToken from "./useRefreshToken";

const usePrivateAxios = () => {
  const refresh = useRefreshToken();
  const { authState, authDispatch } = useContext(authContext);

  useEffect(() => {
    // private route へのアクセス時に access token を付与
    const requestInterceptor = axiosInstance.interceptors.request.use(
      (config) => {
        if (!config.headers["Authorization"]) {
          config.headers["Authorization"] = `Bearer ${authState.accessToken}`;
        }
        return config;
      },
      (error) => Promise.reject(error)
    );

    // private route へのアクセス時にエラーが発生した場合
    const responseInterceptor = axiosInstance.interceptors.response.use(
      (response) => response,
      async (error) => {
        const prevRequest = error.config;
        // access token の有効期限が切れている場合
        if (prevRequest?.url !== "api/auth/login" && error.response) {
          if (error.response.status === 401 && !prevRequest._retry) {
            prevRequest._retry = true;
            // refresh token から新しい access token を取得
            try {
              const newAccessToken = await refresh();
              authDispatch({
                type: AuthActionEnum.SET_ACCESS_TOKEN,
                payload: newAccessToken,
              });
              prevRequest.headers["Authorization"] = `Bearer ${newAccessToken}`;
              return axiosInstance(prevRequest);
            } catch (error) {
              // TODO: session が切れている場合 (refresh token が有効期限切れの場合) -> 再度ログイン必要
              console.log(error);
              return Promise.reject(error);
            }
          }
        }

        return Promise.reject(error);
      }
    );

    return () => {
      axiosInstance.interceptors.response.eject(responseInterceptor);
      axiosInstance.interceptors.request.eject(requestInterceptor);
    };
  }, [authState, authDispatch, refresh]);

  return axiosInstance;
};

export default usePrivateAxios;
