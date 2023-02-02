import { VStack, Heading, Select, HStack } from 'native-base'
import { range } from '../../../../libs/range';
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';

export const BirthDateForm = () => {
  const {onPressBackButton, onPressSubmitButton} = useRegisterProfileContext()

  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>生年月日</Heading>
        <HStack space={2} >
          <Select placeholder="年" size="xl" width="110">
            {
              range(1900,2010).reverse().map((year) => {
                return  <Select.Item key={`${year}`}label={`${year}`} value={`${year}`}/>
              })
            }
          </Select>
          <Select placeholder="月" size="xl" width="100">
            {
              range(1,13).map((month) => {
                return <Select.Item key={`${month}`} label={`${month}`} value={`${month}`}/>
              })
            }
          </Select>
          <Select placeholder="日" size="xl" width="100">
            {
              range(1,32).map((day) => {
                return <Select.Item key={`${day}`} label={`${day}`} value={`${day}`}/>
              })
            }
          </Select>
        </HStack>
        <SubmitButton onPress={onPressSubmitButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
