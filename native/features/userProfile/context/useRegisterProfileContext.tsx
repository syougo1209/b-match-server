import { useContext } from "react"
import { RegisterProfileContext } from "./RegisterProfileContext"

export const useRegisterProfileContext = () => {
  return useContext(RegisterProfileContext)
}
