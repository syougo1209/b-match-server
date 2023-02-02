import { VStack, Heading } from 'native-base'
import { CustomInput } from '../../../../components/elements/CustomInput';
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';
import { useController, useFormContext, RefCallBack } from 'react-hook-form';
import { FC, ChangeEvent } from 'react'

type TNickNameFormProps = {
  onPressBackButton: ()=>void
  onPressSubmitButton: ()=> void
  onChangeText: ((text: string) => void)
  onBlur: () => void
  inputRef: RefCallBack
  fieldValue: string
}

export const NickNameForm: FC<TNickNameFormProps> = (props) => {
  const {onPressBackButton, onPressSubmitButton, onChangeText, onBlur, fieldValue, inputRef } = props

  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>ニックネーム</Heading>
        <CustomInput width="100%" placeholder="nickname" onChangeText={onChangeText} onBlur={onBlur} value={fieldValue} ref={inputRef}/>
        <SubmitButton onPress={onPressSubmitButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
