type User = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  address: string;
  role_id?: number;
  is_admin: bool;
  created_at: string;
};

export type UserLoginResponse = {
  session_id: string;
  access_token: string;
  access_token_expires_at: string;
  refresh_token: string;
  refresh_token_expired_at: string;
  user: User;
};

export type RefreshAccessTokenResponse = {
  access_token: string;
  access_token_expires_at: string;
};
