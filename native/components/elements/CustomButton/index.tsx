import { Button } from 'native-base'
import { FC, ReactNode } from 'react'

type CustomButtonProps = {
  children: ReactNode
}

export const CustomButton: FC<CustomButtonProps> = (props) => {
  return (
    <Button size='lg' borderRadius="50%" fontWeight='bold' _text={{fontSize: "2xl", fontWeight: "bold"}} bgColor='yellow.300'>{props.children}</Button>
  )
}
