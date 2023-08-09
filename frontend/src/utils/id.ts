import { customAlphabet } from "nanoid";

export const generateId = () => {
  return customAlphabet(
    "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz-",
    12
  )();
};
