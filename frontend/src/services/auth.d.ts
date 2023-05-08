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
  access_token: string;
  user: User;
};
