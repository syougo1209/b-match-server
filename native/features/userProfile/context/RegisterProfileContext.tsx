import { createContext } from "react";

type RegisterProfileContextProps = {
  onPressBackButton: () => void
  onPressSubmitButton: () => void
  step: number
}
export const RegisterProfileContext = createContext<RegisterProfileContextProps>({} as RegisterProfileContextProps)
