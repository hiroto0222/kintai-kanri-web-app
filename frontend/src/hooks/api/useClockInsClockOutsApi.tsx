import { AxiosError, AxiosInstance } from "axios";
import { useContext } from "react";
import { toast } from "react-hot-toast";
import { useTranslation } from "react-i18next";
import { authContext } from "../../context/auth";

const useClockInsClockOutsApi = (privateApi: AxiosInstance) => {
  const { t } = useTranslation();
  const { authState } = useContext(authContext);

  const clockIn = async () => {
    const url = "clockins";
    try {
      await privateApi.post(
        url,
        {
          employee_id: authState.user?.id,
        },
        {
          withCredentials: true,
        }
      );
      toast.success(t("success.clock_in"));
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.status === 400) {
          toast.error(t("errors.invalidClockIn"));
          return;
        } else if (error.response?.status === 401) {
          toast.error(t("errors.unauthorized"));
          return;
        }
      }
      toast.error(`an error occured, ${error}`);
    }
  };

  const clockOut = async () => {
    const url = "clockouts";
    try {
      await privateApi.post(
        url,
        { employee_id: authState.user?.id },
        { withCredentials: true }
      );
      toast.success(t("success.clock_out"));
    } catch (error) {
      if (error instanceof AxiosError) {
        if (error.response?.status === 400) {
          toast.error(t("errors.invalidClockOut"));
          return;
        } else if (error.response?.status === 401) {
          toast.error(t("errors.unauthorized"));
          return;
        }
      }
      toast.error(`an error occured, ${error}`);
    }
  };

  return { clockIn, clockOut };
};

export default useClockInsClockOutsApi;
