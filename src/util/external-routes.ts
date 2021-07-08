import { FRONTEND_BASE } from '../../auth-tokens.json';

export const frontend_base = FRONTEND_BASE;

export const frontend = {
  home: frontend_base,
  google: `${frontend_base}auth/google`,
  twitter: `${frontend_base}auth/twitter`,
}