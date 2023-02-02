import { VStack, Heading, Radio, Text, Box } from 'native-base'
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';

export const SexTypeForm = () => {
  const {onPressBackButton, onPressSubmitButton} = useRegisterProfileContext()
  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>性別</Heading>
        <Box ml={3}>
          <Radio.Group name="sexType">
            <Radio value="男性" colorScheme="black">
              <Text fontSize="2xl"  fontWeight="bold">男性</Text>
            </Radio>
            <Radio value="女性" my={5} colorScheme="black">
              <Text fontSize="2xl" fontWeight="bold">女性</Text>
            </Radio>
          </Radio.Group>
        </Box>
        <SubmitButton onPress={onPressSubmitButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
