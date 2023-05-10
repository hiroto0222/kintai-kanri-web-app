export type Employee = {
  id: string;
  first_name: string;
  last_name: string;
  email: string;
  phone: string;
  address: string;
  role_id: {
    int32: number; // sql.nullint32
    valid: boolean;
  };
  is_admin: bool;
  created_at: string;
};
