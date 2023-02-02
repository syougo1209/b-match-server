import { Container } from 'native-base'
import { MailAddressForm } from '../../features/userProfile/components/MailAddressForm';
import { ReactNode, useState } from 'react'
import { BirthDateForm } from '../../features/userProfile/components/BirthDateForm';
import { PasswordForm } from '../../features/userProfile/components/PasswordForm'
import { RegisterProfileContext } from '../../features/userProfile/context/RegisterProfileContext';
import { PrefectureForm } from '../../features/userProfile/components/PrefectureForm';
import { SexTypeForm } from '../../features/userProfile/components/SexTypeForm'
import { NickNameForm } from '../../features/userProfile/components/NickNameForm'

const RegisterComponents: ReactNode[] = [
  <MailAddressForm />,
  <PasswordForm />,
  <NickNameForm />,
  <SexTypeForm />,
  <BirthDateForm />,
  <PrefectureForm />,
]

export const RegistrationScreen = () => {
  const [step, setStep] = useState<number>(0)
  const onPressBackButton = () => setStep((step)=> step - 1)
  const onPressNextButton = () => setStep((step)=> step + 1)

  return (
    <Container safeArea maxW="100%" bgColor="white" flex={1}>
      <RegisterProfileContext.Provider value={{step, onPressBackButton, onPressNextButton}}>
        {RegisterComponents[step]}
      </RegisterProfileContext.Provider>
    </Container>
  );
};
