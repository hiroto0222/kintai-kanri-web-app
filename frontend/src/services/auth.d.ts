import { Employee } from "./employees";

export type UserLoginResponse = {
  session_id: string;
  access_token: string;
  access_token_expires_at: string;
  refresh_token: string;
  refresh_token_expired_at: string;
  user: Employee;
};

export type RefreshAccessTokenResponse = {
  access_token: string;
  access_token_expires_at: string;
};
  