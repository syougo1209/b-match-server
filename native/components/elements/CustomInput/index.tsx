import { Input } from 'native-base'
import { FC } from 'react'

type CustomInputProps = {
  width?: string
  placeholder?: string 
}

export const CustomInput: FC<CustomInputProps> =(props) => {
  return (
    <Input variant="filled" _focus={{bg: "white"}} size="2xl" borderRadius="xl" borderWidth={2} h="16" focusOutlineColor="#facc15" {...props} />
  )
}
