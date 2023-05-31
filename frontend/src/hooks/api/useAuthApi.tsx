import { AxiosError, AxiosInstance, AxiosResponse } from "axios";
import { useContext } from "react";
import { toast } from "react-hot-toast";
import { useTranslation } from "react-i18next";
import { RegisterFormProps } from "../../components/auth/RegisterForm";
import { authContext } from "../../context/auth";
import { AuthActionEnum } from "../../context/auth/authAction";
import api from "../../services/api";
import { UserLoginResponse } from "../../services/auth";

const useAuthApi = (privateApi: AxiosInstance) => {
  const { t } = useTranslation();
  const { authDispatch } = useContext(authContext);

  const login = async (email: string, password: string) => {
    const url = "auth/login";
    try {
      const response: AxiosResponse<UserLoginResponse> = await api.post(url, {
        email,
        password,
      });

      authDispatch({
        type: AuthActionEnum.LOG_IN,
        payload: response.data,
      });
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.status === 401) {
          toast.error(t("errors.invalidUserCredentials"));
          return;
        } else if (error.response?.status === 404) {
          toast.error(t("errors.userNotFound"));
          return;
        }
      }
      toast.error(`an error occured, ${error}`);
    }
  };

  const register = async (data: RegisterFormProps) => {
    const url = "auth/register";
    try {
      await privateApi.post(
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
      toast.success(t("success.registered"));
    } catch (error) {
      toast.error(`an error occured, ${error}`);
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
