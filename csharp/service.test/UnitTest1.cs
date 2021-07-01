using System;
using Xunit;
using Service;
using FluentValidation;


namespace Service.Test
{
    public class UnitTest1
    {
        [Fact]
        public void Test1()
        {
            Person p = new Person();
            PersonValidator pv = new PersonValidator();

            // ValidationResult result = pv.Validate(p);

            Console.Write(pv.Validate(p));
        }
    }
}
