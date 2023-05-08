import { AxiosResponse } from "axios";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { authContext } from "../../context/auth";
import { AuthActionEnum } from "../../context/auth/authAction";
import api from "../../services/api";
import { UserLoginResponse } from "../../services/auth";

const useAuthApi = () => {
  const { authDispatch } = useContext(authContext);
  const navigate = useNavigate();

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
      console.log(response.data);
      navigate("/me");
    } catch (error) {
      console.log(error);
    }
  };

  const logout = () => {
    authDispatch({
      type: AuthActionEnum.LOG_OUT,
    });
  };

  return { login, logout };
};

export default useAuthApi;
