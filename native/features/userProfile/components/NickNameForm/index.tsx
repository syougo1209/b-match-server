import { VStack, Heading, Box } from 'native-base'
import { CustomInput } from '../../../../components/elements/CustomInput';
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';

export const NickNameForm = () => {
  const {onPressBackButton, onPressSubmitButton} = useRegisterProfileContext()
  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>ニックネーム</Heading>
        <CustomInput width="100%" placeholder="nickname"/>
        <SubmitButton onPress={onPressSubmitButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
