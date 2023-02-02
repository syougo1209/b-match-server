import { createContext } from "react";

type RegisterProfileContextProps = {
  onPressBackButton: () => void
  onPressNextButton: () => void
  step: number
}
export const RegisterProfileContext = createContext<RegisterProfileContextProps>({} as RegisterProfileContextProps)
