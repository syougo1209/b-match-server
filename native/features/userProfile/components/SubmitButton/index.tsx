import { CustomButton } from '../../../../components/elements/CustomButton';
import { FC, ReactNode } from 'react'

type SubmitButtonProps = {
  onPress: () => void
  children: ReactNode
}
export const SubmitButton: FC<SubmitButtonProps> = (props) => {
  return <CustomButton onPress={props.onPress}>{props.children}</CustomButton>
}
