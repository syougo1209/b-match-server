import { VStack, Heading, Box } from 'native-base'
import { CustomInput } from '../../../../components/elements/CustomInput';
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';

export const PasswordForm = () => {
  const {onPressBackButton, onPressNextButton} = useRegisterProfileContext()
  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>パスワード</Heading>
        <CustomInput width="100%" placeholder="lg"/>
        <SubmitButton onPress={onPressNextButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
