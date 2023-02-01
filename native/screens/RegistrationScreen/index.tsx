import { useNavigation} from '@react-navigation/native'
import { NativeStackNavigationProp} from '@react-navigation/native-stack';
import { Screens } from '../../Screens'
import { Container, Box, Center, Icon, IconButton } from 'native-base'
import { CustomInput } from '../../components/elements/CustomInput'
import { ChevronIconButton } from '../../components/elements/ChevronIconButton';
export const RegistrationScreen = () => {
  const navigation = useNavigation<NativeStackNavigationProp<Screens, 'Conversations'>>()

  return (
    <Container safeArea maxW="100%" bgColor="white">
      <ChevronIconButton onPress={()=>navigation.navigate('Conversations')}/>
      <Center width="100%">
        <Box alignItems="center">
          <CustomInput width="80%" placeholder="lg"/>
        </Box>
        <Box>hohoge</Box>
      </Center>
    </Container>
  );
};
