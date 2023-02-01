import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { Container, VStack, Heading } from 'native-base'
import { CustomInput } from '../../components/elements/CustomInput'
import { ChevronIconButton } from '../../components/elements/ChevronIconButton';
import { CustomButton } from '../../components/elements/CustomButton';

export const RegistrationScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()

  return (
    <Container safeArea maxW="100%" bgColor="white" flex={1}>
      <ChevronIconButton onPress={()=>navigation.navigate('Conversations')}/>
      <VStack space={10} width="100%" p={4}>
        <Heading size='xl'>メールアドレス</Heading>
        <CustomInput width="100%" placeholder="lg"/>
        <CustomButton>次へ</CustomButton>
      </VStack>
    </Container>
  );
};
