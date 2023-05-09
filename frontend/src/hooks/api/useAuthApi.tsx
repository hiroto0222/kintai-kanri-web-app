import { AxiosResponse } from "axios";
import { useContext } from "react";
import { useNavigate } from "react-router-dom";
import { RegisterFormProps } from "../../components/auth/RegisterForm";
import { authContext } from "../../context/auth";
import { AuthActionEnum } from "../../context/auth/authAction";
import api from "../../services/api";
import { UserLoginResponse } from "../../services/auth";
import usePrivateAxios from "../usePrivateAxios";

const useAuthApi = () => {
  const privateApi = usePrivateAxios();
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

  const register = async (data: RegisterFormProps) => {
    const url = "auth/register";
    try {
      const response = await privateApi.post(
        url,
        {
          first_name: data.FirstName,
          last_name: data.LastName,
          email: data.Email,
          password: data.Password,
          phone: data.Phone,
          address: data.Address,
          role_id: data.RoleID,
          is_admin: data.IsAdmin,
        },
        {
          withCredentials: true,
        }
      );
      console.log(response.data);
    } catch (error) {
      console.log(error);
    }
  };

  const logout = () => {
    authDispatch({
      type: AuthActionEnum.LOG_OUT,
    });
  };

  return { login, logout, register };
};

export default useAuthApi;