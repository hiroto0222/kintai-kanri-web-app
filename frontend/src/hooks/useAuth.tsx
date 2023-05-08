import { AxiosResponse } from "axios";
import { useContext } from "react";
import { authContext } from "../context/auth";
import { AuthActionEnum } from "../context/auth/authAction";
import api from "../services/api";
import { UserLoginResponse } from "../services/auth";

const useAuth = () => {
  const { authDispatch } = useContext(authContext);

  const login = async (email: string, password: string) => {
    const url = "auth/login";
    try {
      const response: AxiosResponse<UserLoginResponse> = await api.post(
        url,
        {
          email,
          password,
        },
        { withCredentials: true }
      );

      authDispatch({
        type: AuthActionEnum.LOG_IN,
        payload: response.data,
      });
    } catch (error) {
      console.log(error);
    }
  };

  return { login };
};

export default useAuth;
