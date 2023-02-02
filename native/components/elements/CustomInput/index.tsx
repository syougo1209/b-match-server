import { Input, IInputProps } from 'native-base'
import { FC } from 'react'


export const CustomInput: FC<IInputProps> =(props) => {
  return (
    <Input variant="filled" _focus={{bg: "white"}} size="2xl" borderRadius="xl" borderWidth={2} h="16" focusOutlineColor="#facc15" {...props} />
  )
}
