import { VStack, Heading, Box, Radio, Text } from 'native-base'
import { ChevronIconButton } from '../../../../components/elements/ChevronIconButton';
import { SubmitButton } from '../../../../features/userProfile/components/SubmitButton';
import { useRegisterProfileContext } from '../../context/useRegisterProfileContext';

export const PrefectureForm = () => {
  const {onPressBackButton, onPressNextButton} = useRegisterProfileContext()
  const prefectures = [{id: 1, name: "北海道"}, {id: 2, name: "秋田"}, {id: 3, name: "宮城"}, {id: 4, name: "岩手"}, {id: 5, name: "福島"}]
  return (
    <>
      <ChevronIconButton onPress={onPressBackButton}/>
      <VStack space={10} width="100%" p={8}>
        <Heading size='xl'>都道府県</Heading>
        <Box ml={3}>
          <Radio.Group name="sexType">
            {
              prefectures.map((prefecture)=>{
                return (
                  <Radio key={`${prefecture.id}`} value={`${prefecture.id}`} colorScheme="black">
                    <Text fontSize="2xl"  fontWeight="bold">{prefecture.name}</Text>
                  </Radio>
                )
            })
            }
          </Radio.Group>
        </Box>
        <SubmitButton onPress={onPressNextButton}>次へ</SubmitButton>
      </VStack>
    </>
  )
}
