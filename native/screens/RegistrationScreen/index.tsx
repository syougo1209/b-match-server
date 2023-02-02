import { Container } from 'native-base'
import { MailAddressForm } from '../../features/userProfile/components/MailAddressForm';
import { ReactNode, useState } from 'react'
import { BirthDateForm } from '../../features/userProfile/components/BirthDateForm';
import { PasswordForm } from '../../features/userProfile/components/PasswordForm'
import { RegisterProfileContext } from '../../features/userProfile/context/RegisterProfileContext';
import { PrefectureForm } from '../../features/userProfile/components/PrefectureForm';
import { SexTypeForm } from '../../features/userProfile/components/SexTypeForm'
import { NickNameFormContainer } from '../../features/userProfile/register/components/NickNameFormContainer';
import { useRegistrationScreen } from './hooks/useRegistrationScreen';
import { useRegisterForm } from '../../features/userProfile/register/hooks/useRegisterForm'
import { FormProvider } from 'react-hook-form'

export const RegisterComponents: ReactNode[] = [
  <MailAddressForm />,
  <PasswordForm />,
  <NickNameFormContainer />,
  <SexTypeForm />,
  <BirthDateForm />,
  <PrefectureForm />,
]

export const RegistrationScreen = () => {
  const {methods, onSubmit } = useRegisterForm()
  const { step, onPressBackButton, onPressSubmitButton } = useRegistrationScreen(onSubmit)

  return (
    <Container safeArea maxW="100%" bgColor="white" flex={1}>
      <FormProvider {...methods}>
        <RegisterProfileContext.Provider value={{step, onPressBackButton, onPressSubmitButton}}>
          {RegisterComponents[step]}
        </RegisterProfileContext.Provider>
      </FormProvider>
    </Container>
  );
};
